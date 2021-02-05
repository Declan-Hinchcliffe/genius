import React from 'react'

function SongCard(props) { 
  
    console.log("props", props.song.title)
    console.log("props", props.song.artist)

    if (props.song.lyrics.lyrics === "") {
        props.song.lyrics.lyrics = "couldn't find lyrics for this song"
    }

    return (
        <div>
            <div>
                <div>
                    { props.song.artist } - { props.song.title }
                </div>
                <div >
                    { props.song.lyrics.lyrics }
                </div>
            </div>
        </div>
    )
}

export default SongCard