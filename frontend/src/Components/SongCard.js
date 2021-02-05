import React from 'react'

function SongCard(props) { 
    if (props.data.lyrics.lyrics === "") {
        props.data.lyrics.lyrics = "couldn't find lyrics for this song"
    }

    return (
        <div>
            <div>
                { props.data.artist } - { props.data.title }
            </div>
            <div >
                { props.data.lyrics.lyrics }
            </div>
        </div>
    )
}

export default SongCard