import { useEffect, useState } from 'react'

export interface PostsWithUsername {
    id: number
    user_id: string
    username: string
    content: string
    image_url: string
    created_at: string
    updated_at: string

}

const useGetAllPosts = () => {
    const [posts, setPosts] = useState<PostsWithUsername[]>([])

    const token = sessionStorage.getItem("Token")

    useEffect(() => {
        fetch(`http://localhost:8000/api/v1/posts?page=1&limit=20`, {
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