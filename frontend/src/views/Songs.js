import React, { useState, useEffect } from 'react'
import Axios from 'axios'
import { useParams } from 'react-router-dom'
import SongCard from '../Components/SongCard'

function Songs() {
    const { name } = useParams()
    const url = `http://localhost:9000/songs/lyrics/${name}`
    const [songs, setSongs] = useState(null)


    let content = null

    useEffect(() => {
        Axios.get(url)
        .then(response => {
            setSongs(response.data)
        })
        .catch(() => {
            
        })
    }, [url])

    console.log(songs)

    if (songs) {
        content = 
        songs.map((song, key) => 
            <div>
                <SongCard
                    song={song}
                />
            </div>
        )

        return (
            <div>
               { content }
            </div>
        )
    }


    return (
        <div>
            
        </div>
    )
}

export default Songs