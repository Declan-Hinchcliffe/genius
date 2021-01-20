import React, { useState, useEffect } from 'react'
import Axios from 'axios'

function Songs() {
    const url = 'http://localhost:9000/lyrics/black'
    const [song, setSong] = useState(null)

    useEffect(() => {
        Axios.get(url)
        .then(response => {
            setSong(response.data)
        })
    }, [url])

    console.log(song)

    if (song) {
        return (
            <div>
                <h1> {song.title} </h1>
                <p> {song.artist} </p>
                <p> {song.lyrics.lyrics} </p>
            </div>
        )
    }


    return (
        <div>

        </div>
    )
}

export default Songs