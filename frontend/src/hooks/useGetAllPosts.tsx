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

    const token = sessionStorage.getItem("Token")
    const HOST = import.meta.env.VITE_API_HOST

    useEffect(() => {
        fetch(`${HOST}/api/v1/posts?page=1&limit=20`, {
            headers: {
                Authorization: `Bearer ${token}`,
            }
        })
            .then(res => res.json())
            .then(result => setPosts(result.data))
            .catch(err => console.error(err))
    }, [])

    return { posts, setPosts }
}


export default useGetAllPosts