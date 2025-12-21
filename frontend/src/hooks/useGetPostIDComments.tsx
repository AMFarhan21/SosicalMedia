import { useEffect, useState } from 'react'
import type { PostsWithUsername } from './useGetAllPosts'

export interface CommentsWithUsername {
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
}

const useGetPostIDComments = (postID: number) => {
    const initialPost: PostsWithUsername = {
        id: 0,
        user_id: "",
        first_name: "",
        last_name: "",
        username: "",
        content: "",
        image_url: [],
        created_at: "",
        updated_at: "",
        is_liked: false,
        likes_count: 0,
        comments_count: 0,
    }

    const [error, setError] = useState("")
    const [comments, setComments] = useState<CommentsWithUsername[]>([])
    const [post, setPost] = useState<PostsWithUsername>(initialPost)
    const token = sessionStorage.getItem("Token")
    const HOST = import.meta.env.VITE_API_HOST


    useEffect(() => {
        const fetchPost = async () => {
            try {
                const res = await fetch(`${HOST}/api/v1/posts/${postID}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const result = await res.json()

                if (!res.ok) {
                    throw new Error(result.message)
                }

                setPost(result.data)
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message)
                } else {
                    setError("Error on fetching post by id")
                }
            }
        }
        const fetchComments = async () => {
            try {
                const res = await fetch(`${HOST}/api/v1/posts/${postID}/comments`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })
                const result = await res.json()

                if (!res.ok) {
                    throw new Error(result.message)
                }


                setComments(result.data)
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message)
                } else {
                    setError("Error on fetching comments")
                }
            }
        }

        fetchPost()
        fetchComments()
    }, [postID, token, HOST])




    return { post, setPost, comments, setComments, error }

}

export default useGetPostIDComments