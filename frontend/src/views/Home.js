import React, { useEffect, useState } from 'react'
import Axios from 'axios'

function Home() {
    const [count, setCount] = useState(0);

    useEffect(() => {
     
    }, [count])

    Axios.post('http://localhost:9000/song/lyrics', {
        name: "wap"
    })
    .then((response) => {
        console.log(response);
    })


    return (
        <div>
            <form>
                <label>Search: </label>
                <input className="border rounded-lg"/>
                <button type="submit" className="w-20 border rounded-lg" onClick={() => setCount(count + 1)}>Submit</button>
            </form>
        </div>
    )
}

export default Home