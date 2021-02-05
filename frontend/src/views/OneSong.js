import React, { useState } from 'react'
// import { useParams } from 'react-router-dom'
// import SongCard from '../Components/SongCard'
// import Loader from '../Components/Loader'
import AxiosGet from '../Components/HttpRequest'
import SongCard from '../Components/SongCard';

function OneSong() {
    const [bool, setBool] = useState(null);
    const [search, setSearch] = useState('');
    const [words, setWords] = useState('');

    const url = 'http://localhost:9000/song/lyrics'

    let content = null
 
    let songs = AxiosGet(url, search, words, bool)

    if (songs.error) {
        content = 
        <div> there was an error, please try again </div>
    }

    if(songs.loading){
        content =
        <div>loading...</div>
    }
    
    if (songs.data){
        content = 
            songs.data.songs.map((song) => {
            let props = {
                song: song,
                wordMap: songs.data.wordMap
            }

        return (    
            <div key={song.id}>
                <SongCard 
                    data={props}
                />
            </div>
            )
        })
    } 
    

    return (
        <div>
            <div>
                <h2>Search for a single song...</h2>
                <form>
                    Search:<input onChange={e => {setSearch(e.target.value)}}></input>
                    Words:<input onChange={e => {setWords(e.target.value)}}></input>
                    <button onClick={() => setBool(!bool)}  type="button">Submit</button>
                </form>

                { content }
            </div>  
        </div>
    )
}

export default OneSong