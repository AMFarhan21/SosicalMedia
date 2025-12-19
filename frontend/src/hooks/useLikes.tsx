import type { PostsWithUsername } from './useGetAllPosts';
import type { CommentsWithUsername } from './useGetPostIDComments';

const useLikes = ({ setPosts, setPost, setComments }: { setPosts?: React.Dispatch<React.SetStateAction<PostsWithUsername[]>>, setPost?: React.Dispatch<React.SetStateAction<PostsWithUsername>>, setComments?: React.Dispatch<React.SetStateAction<CommentsWithUsername[]>> }) => {
    const token = sessionStorage.getItem("Token")
    const HOST = import.meta.env.VITE_API_HOST;

    const likes = async (target_id: number, target: "POST" | "COMMENT") => {
        try {
            let body = {}
            if (target == "POST") {
                body = { post_id: target_id }
            } else {
                body = { comment_id: target_id }
            }

            const res = await fetch(`${HOST}/api/v1/likes`, {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${token}`,
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(body)
            })

            const result = await res?.json()

            if (!res.ok) {
                throw new Error(result.error)
            }

            let returning
            if (result.data == "Unliked") {
                returning = false
            } else {
                returning = true
            }

            if (setPosts) {
                setPosts((prev) => prev.map((post) => post.id == target_id ? { ...post, is_liked: returning, likes_count: returning ? post.likes_count + 1 : post.likes_count - 1 } : post))
            } else if (setPost) {
                setPost((prev: PostsWithUsername) => ({ ...prev, is_liked: returning, likes_count: returning ? prev.likes_count + 1 : prev.likes_count - 1 }))
            } else if (setComments) {
                setComments((prev) => prev.map((post) => post.id == target_id ? { ...post, is_liked: returning, likes_count: returning ? post.likes_count + 1 : post.likes_count - 1 } : post))
            }

            return returning

        } catch (err) {
            if (err instanceof Error) {
                console.error(err.message)
            } else {
                console.error("Likes failed, please try again!")
            }
        }

    }


    return { likes }

}

export default useLikes