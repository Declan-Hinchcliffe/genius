import { useEffect, useState } from "react"
import axios from 'axios'

function AxiosGet(url, search, words, bool) {
    const [request, setRequest] = useState(null)

    // let content = null

       useEffect(() => {
        axios.interceptors.request.use(request => {
            console.log(JSON.stringify(request.data))
            return request
        })

        axios({
            method: 'post',
            url: url,
            data: { 
                search: search,
                words: words
            }
        })
        .then((response) => {
            console.log(response.data);
            setRequest(response.data);
        })
    }, [bool])

    // useEffect(() => {
    //     setRequest({
    //         loading: true,
    //         data: null,
    //         error: false,
    //     })
    //     Axios.get(url)
    //     .then(response => {
    //         setRequest({
    //             loading: false,
    //             data: response.data,
    //             error: false,
    //         })
    //     })
    //     .catch(() => {
    //         setRequest({
    //             loading: false,
    //             data: null,
    //             error: true,
    //         })
    //     })
    // }, [url])

    return request
}

export default AxiosGet