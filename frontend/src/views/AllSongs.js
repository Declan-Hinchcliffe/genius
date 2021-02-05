import React, { useState } from 'react'
import AxiosGet from '../Components/HttpRequest';
import SongCard from '../Components/SongCard';

function AllSongs() {
    const [bool, setBool] = useState(null);
    const [search, setSearch] = useState('');
    const [words, setWords] = useState('');
 
    const url = 'http://localhost:9000/songs/lyrics/artist'
    
    let songs = AxiosGet(url, search, words, bool)

    let songData = null
    let wordsObj = null
    let wordsCounter = null

    if (songs.error) {
        songData = 
        <div> there was an error, please try again </div>
    }

    if(songs.loading){
        songData =
        <div>loading...</div>
    }
    
    if (songs.data){
        wordsObj = songs.data.wordMap
        wordsCounter =               
        <div>
            { 
                Object.entries(wordsObj).map(([key,value], index) => {
                    return (
                        <div key={index}>{key}: {value.toString()}</div>
                    );
                })
            }
        </div>

        songData = 
            songs.data.songs.map((song) => {

        return (    
            <div key={song.id}>
                <SongCard 
                    data={song}
                />
            </div>     
            )
        })
    }

    return (
        <div>
            <div>
                <h2>Matches the top 20 songs by an artist...</h2>
                <form>
                    <label>Search: </label>
                        <input className="border rounded-lg" onChange={e => {setSearch(e.target.value)}}></input>
                    <label>Words: </label>
                        <input className="border rounded-lg" onChange={e => {setWords(e.target.value)}}></input>
                    <button onClick={() => setBool(!bool)}  type="button" className="w-20 border rounded-lg">Submit</button>
                </form>
                <div>
                    { wordsCounter }
                </div>
                <div>
                    { songData }
                </div>
            </div>
        </div>
    )
}

export default AllSongs;