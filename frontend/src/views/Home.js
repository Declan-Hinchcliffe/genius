import React, { useEffect, useState } from 'react'
import axios from 'axios'

function Home() { 
    const [foo, setFoo] = useState(false);
 
   useEffect(() => {
        axios.interceptors.request.use(request => {
            console.log(JSON.stringify(request))
            return request
        })

        axios({
            method: 'post',
            url: 'http://localhost:9000/song/lyrics',
            data: { 
                name: 'hello'
            }
        })
        .then((response) => {
            console.log(response.data);
        })
    }, [foo])

    return (
        <div>
            <form>
                <label>Search: </label>
                <input className="border rounded-lg"/>
                <button onClick={()=> {setFoo(!foo)}}  type="button" className="w-20 border rounded-lg">Submit</button>
            </form>
        </div>
    )
}

export default Home