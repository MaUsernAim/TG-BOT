// Находим адрес функции по имени
const targetFunc = "main.processData"; // ЗАМЕНИТЕ на ваше имя функции
const addr = DebugSymbol.fromName(targetFunc).address;

Interceptor.attach(addr, {
    onEnter: function (args) {
        // В Go 1.22 на amd64:
        // RAX = указатель на строку
        // RBX = длина строки
        let strPtr = this.context.rax;
        let strLen = this.context.rbx.toInt32();

        if (!strPtr.isNull() && strLen > 0) {
            // 1. Читаем текущее значение
            let currentStr = Memory.readUtf8String(strPtr, strLen);
            console.log("[*] Оригинальная строка: " + currentStr);

            // 2. Меняем значение на "HACKED"
            let newStr = "HACKED";
            
            // ВАЖНО: Мы не можем записать строку длиннее, чем оригинал, 
            // не выходя за пределы выделенной памяти. 
            // Если новая строка короче — это безопасно.
            if (newStr.length <= strLen) {
                Memory.writeUtf8String(strPtr, newStr);
                
                // Если новая строка короче, обновляем длину в регистре RBX
                this.context.rbx = ptr(newStr.length);
                
                console.log("[+] Значение успешно изменено!");
            } else {
                console.log("[!] Ошибка: Новая строка длиннее оригинальной. Риск падения программы.");
            }
        }
    }
});