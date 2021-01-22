import React from 'react'

function SongCard(props) {
    if (props.song.lyrics.lyrics === "") {
        props.song.lyrics.lyrics = "Unable to find lyrics for this song!"
    }

    return (
        <div className="p-5">
            <div className="w-full border mb-4 rounded-lg overflow-hidden">
                <div className="p-2 border text-sm">
                    { props.song.artist } - { props.song.title }
                </div>
                <div className="p-2 border text-sm">
                    { props.song.lyrics.lyrics }
                </div>
            </div>
        </div>
    )
}

export default SongCard