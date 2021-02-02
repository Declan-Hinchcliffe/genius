import React, { useEffect, useState } from 'react'
import axios from 'axios'

function Home() { 
    const [bool, setBool] = useState(null);
    const [search, setSearch] = useState('');
    const [words, setWords] = useState('');
 
   useEffect(() => {
        axios.interceptors.request.use(request => {
            console.log(JSON.stringify(request.data))
            return request
        })

        axios({
            method: 'post',
            url: 'http://localhost:9000/songs/lyrics/artist',
            data: { 
                search: `${search}`,
                words:`${words}`
            }
        })
        .then((response) => {
            console.log(response.data);
        })
    }, [bool])

    return (
        <div>
            <form>
                <label>Search: </label>
                <input className="border rounded-lg" onChange={e => {setSearch(e.target.value)}}></input>
                <input className="border rounded-lg" onChange={e => {setWords(e.target.value)}}></input>
                <button onClick={() => setBool(!bool)}  type="button" className="w-20 border rounded-lg">Submit</button>
            </form>
        </div>
    )
}

export default Home