import { useEffect, useState } from "react"
import axios from 'axios'

function AxiosGet(url, search, words, bool) {
    const [request, setRequest] = useState({
        loading: false,
        data: null,
        error: false
    })

    useEffect(() => {
        axios.interceptors.request.use(request => {
            console.log(JSON.stringify(request.data))
            return request
        })
        setRequest({
            loading: true,
            data: null,
            error: false
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
            setRequest({
                loading: true,
                data: response.data,
                error: false,
            })
            
            console.log(response.data);
        })
        .catch(() => {
            setRequest({
                loading: false,
                data: null,
                error: true,
            })
        })
    }, [bool])

    return request
}

export default AxiosGet