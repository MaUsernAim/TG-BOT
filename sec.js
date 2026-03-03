import React ,{useState} from 'react'

function App(){

    const [iV,sIV]=useState("")
    const [dT,sDT]=useState("")

    const H1=async ()=>{
        const resp= await fetch("http://10.111.89.118:8001/api/",{
          method:"POST",
          headers:{
            "Content-Type":"application/json",
          },
          body:JSON.stringify({
            message:iV
          }),
        });

        const data=await resp.json()
        sDT(data.message)
    } 

    return(
        <div class="main">

            <div class="message_area">

                <h1>{dT}</h1>

            </div>

            <div class="input_area">
            <input 
                class="input_in_area"
                text="text"
                value={iV}
                onChange={(e)=>sIV(e.target.value)}
                />
            
            <button onClick={H1}></button>
            </div>

        </div>
    )

}

export default App;