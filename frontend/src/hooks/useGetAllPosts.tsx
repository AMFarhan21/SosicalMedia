import { useEffect, useState } from 'react'

export interface PostsWithUsername {
    id: number
    user_id: string
    first_name: string
    last_name: string
    username: string
    content: string
    image_url: string[]
    created_at: string
    updated_at: string
    is_liked: boolean
    likes_count: number
    comments_count: number
}

const useGetAllPosts = () => {
    const [posts, setPosts] = useState<PostsWithUsername[]>([])
    const [loadingGetAllPosts, setLoading] = useState(false)
    const token = localStorage.getItem("Token")
    const HOST = import.meta.env.VITE_API_HOST

    useEffect(() => {
        const fetchAllPost = async () => {
            try {
                setLoading(true)
                await new Promise(res => setTimeout(res, 1500)) // TEST
                const res = await fetch(`${HOST}/api/v1/posts?page=1&limit=20`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    }
                })
                const result = await res.json()
                if (!res.ok) {
                    throw new Error(result.error)
                }
                setPosts(result.data)
            } catch (err) {
                console.error(err)
            } finally {
                setLoading(false)
            }
        }

        fetchAllPost()
    }, [token, HOST])

    return { posts, setPosts, loadingGetAllPosts }
}


export default useGetAllPosts