import React from 'react'
import { useParams } from 'react-router-dom'
import SongCard from '../Components/SongCard'
import Loader from '../Components/Loader'
import AxiosGet from '../Components/HttpRequest'

function Songs() {
    const { name } = useParams()
    const url = `http://localhost:9000/songs/lyrics/${name}`
   
   let songs = AxiosGet(url)
    
    let content = null

    if (songs.error) {
        content = 
        <div className="p-5">
            <div className="bg-red-300 mb-2 p-3 rounded-lg">
                There was an error, please refresh or try again.
            </div>
        </div>
    }

    if (songs.loading) {
        content = <Loader></Loader>
    }

    if (songs.data) {
        content = 
        songs.data.map((song) => 
            <div key={song.id}>
                <SongCard
                    song={song}
                />
            </div>
        )
    }

    return (
        <div>
            { content }
        </div>
    )
}

export default Songs