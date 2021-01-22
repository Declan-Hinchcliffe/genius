import React from 'react'

function SongCard(props) {
    return (
        <div className="p-5">
            <div className="w-full border mb-4 rounded-lg overflow-hidden">
                <p className="p-2 border">{ props.song.artist }</p>
                <p className="p-2 border">{ props.song.title }</p>
                <p className="p-2 border">{ props.song.lyrics.lyrics }</p>
            </div>
        </div>
    )
}

export default SongCard