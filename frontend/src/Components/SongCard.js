import React from 'react'

function SongCard(props) {
    if (props.lyrics.lyrics === "") {
        props.lyrics.lyrics = "Unable to find lyrics for this song!"
    }

    return (
        <div>
            <div>
                <div >
                    { props.artist } - { props.title }
                    
                </div>
                <div >
                    { props.lyrics.lyrics }
                </div>
            </div>
        </div>
    )
}

export default SongCard