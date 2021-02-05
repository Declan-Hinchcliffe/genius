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

        const doRequest = async () => {
            setRequest({
                loading: true,
                data: null,
                error: false
            })
            try {
                const response = await axios.post(url, {
                    search,
                    words
                })

                setRequest({
                    loading: true,
                    data: response.data,
                    error: false,
                })
            } catch {
                setRequest({
                    loading: false,
                    data: null,
                    error: true,
                })
            }
        }
        doRequest()

    }, [bool])

    return request
}

export default AxiosGet