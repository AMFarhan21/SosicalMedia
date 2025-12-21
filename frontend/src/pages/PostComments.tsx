import { useParams } from "react-router-dom"
import useGetPostIDComments from "../hooks/useGetPostIDComments"
import { Image, X } from 'lucide-react'
import { useEffect, useState } from 'react'
import Post from "../components/Post"
import Comments from "../components/Comments"
import useCreateComments from "../hooks/useCreateComments"



const PostComments = () => {
    const [commentInput, setCommentInput] = useState("")
    const [img, setImg] = useState("")
    const [showImg, setShowImg] = useState(false)
    const [files, setFiles] = useState<File[]>([])

    const param = useParams()
    const HOST = import.meta.env.VITE_API_HOST

    const { post, setPost, comments, setComments, error } = useGetPostIDComments(Number(param.postID))


    const { createComments, errorComments } = useCreateComments()


    const handleCreateComment = (e: React.FormEvent) => {
        e.preventDefault()
        if (!post) return
        createComments(post.id, commentInput, files, setComments)
        setCommentInput("")
        setFiles([])
    }


    useEffect(() => {
        if (error != "") {
            alert(error)
        }
    }, [error])

    useEffect(() => {
        if (img) {
            const timer = setTimeout(() => setShowImg(true), 10)
            return () => clearTimeout(timer)
        } else {
            const timer = setTimeout(() => setShowImg(false), 10)
            return () => clearTimeout(timer)
        }
    }, [img])




    return (
        <>
            {
                img && (
                    <div className={`fixed inset-0 bg-black/80 z-51 items-center m-auto flex transition-opacity scale-100 duration-300 ${showImg ? "opacity-100" : "opacity-0"}`}>
                        <img src={`${HOST}/${img}`} className={`max-h-full mx-auto transition-opacity scale-100 duration-300 ${showImg ? "scale-100" : "scale-0"}`} />
                        <div onClick={() => setImg("")} className='absolute top-10 right-10 cursor-pointer'>
                            <X />
                        </div>
                    </div>
                )
            }
            <div className="w-full sm:w-[56%] mx-auto">
                {
                    post && <Post post={post} setPost={setPost} idxCondition={true} postCommentsPage={() => { }} onPostComment={true} setImg={setImg} />
                }
                <form onSubmit={handleCreateComment} className={`sm:border-x ${!comments && "border-b"} border-white/22 relative cursor-pointer p-4 flex flex-col space-y-4`}>
                    <textarea value={commentInput} onChange={(e) => setCommentInput(e.target.value)} placeholder={`Reply to @${post.username}`} className="outline-none" />
                    <label htmlFor="upload-image">
                        <div className="cursor-pointer hover:bg-blue-400/20 duration-100 p-2 w-8 rounded-full -mb-6">
                            <Image className="w-4 h-4 text-blue-500" />
                        </div>
                    </label>
                    <input id="upload-image" type="file" accept="image/*" multiple onChange={(e) => {
                        if (!e.target.files) return
                        setFiles(Array.from(e.target.files))
                    }} placeholder={`Image url`} className="hidden outline-none" />
                    {
                        error && (
                            <div className="text-red-500">
                                {errorComments}
                            </div>
                        )
                    }
                    <button className="absolute right-4 bottom-4 bg-white text-black font-bold px-4 py-1 cursor-pointer hover:bg-white/80 duration-100 rounded-full" type="submit">Reply</button>
                </form>
                <div>
                    {
                        comments && comments.map((comment) => (
                            <Comments post_id={post.id} username={post.username} comment={comment} setComments={setComments} />
                        ))
                    }
                </div>
            </div>
        </>
    )
}

export default PostComments