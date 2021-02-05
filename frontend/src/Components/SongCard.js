import React from 'react'

function SongCard(props) { 
    let obj = props.data.wordMap

    if (props.data.song.lyrics.lyrics === "") {
        props.data.song.lyrics.lyrics = "couldn't find lyrics for this song"
    }

    let wordCount = null

    if (obj) {
        wordCount = 
        <div>
            { 
                Object.entries(obj).map(([key,value], index) => {
                    return (
                        <div key={index}>{key}: {value.toString()}</div>
                    );
                })
            }
        </div>
    }

    return (
        <div>
            <div>
                { props.data.song.artist } - { props.data.song.title }
            </div>
            <div >
                { props.data.song.lyrics.lyrics }
            </div>
            <div>
                { wordCount }
            </div>
           
        </div>
    )
}

export default SongCard