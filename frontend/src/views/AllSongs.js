import React, { useEffect, useState } from 'react'
import axios from 'axios'
import AxiosGet from '../Components/HttpRequest';
import SongCard from '../Components/SongCard';

function AllSongs() {
    const [bool, setBool] = useState(null);
    const [search, setSearch] = useState('');
    const [words, setWords] = useState('');
 
    const url = 'http://localhost:9000/songs/lyrics/artist'
    
    let songs = AxiosGet(url, search, words, bool)

    let content = null

     
    if (songs){
     
        content = 
        song.songs.map((song, key) => {
            <div>
                { song.title }
                { song.artist }
                { song.lyrics.lyrics }
            </div>

            console.log(song.title, song.artist)

        })
        
    }

    return (
        <div>
            <div>
                <h2>Search for songs...</h2>
                <form>
                    <label>Search: </label>
                        <input className="border rounded-lg" onChange={e => {setSearch(e.target.value)}}></input>
                    <label>Words: </label>
                        <input className="border rounded-lg" onChange={e => {setWords(e.target.value)}}></input>
                    <button onClick={() => setBool(!bool)}  type="button" className="w-20 border rounded-lg">Submit</button>
                </form>


                { content }
            </div>
        </div>
    )
}

export default AllSongs;