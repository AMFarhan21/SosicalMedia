import { useEffect, useState } from 'react'
import useCreatePosts from '../hooks/useCreatePosts'
import Posts from '../components/Post'
import useGetAllPosts from '../hooks/useGetAllPosts'
import { useNavigate } from 'react-router-dom'
import { Image, X } from 'lucide-react'


const Feed = () => {
    const [content, setContent] = useState("")
    const [files, setFiles] = useState<File[]>([])
    const [img, setImg] = useState("")

    const HOST = import.meta.env.VITE_API_HOST

    const { posts, setPosts } = useGetAllPosts()
    const { createPost, errorCreate } = useCreatePosts(setPosts)
    const navigate = useNavigate()

    const handleCreatePost = (e: React.FormEvent) => {
        e.preventDefault()
        createPost(content, files)


        setContent("")
        setFiles([])
    }

    const [showImg, setShowImg] = useState(false)

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
            <div className='w-full sm:w-[56%] mx-auto'>
                <div className='border-b sm:border border-white/22 font-bold text-lg text-center mx-auto p-3 backdrop-blur-sm sticky top-0 bg-black/80 z-50'>
                    For you
                    <br />
                    <div className='absolute left-0 right-0 top-9 text-blue-500 font-bold'>
                        ---------
                    </div>
                </div>
                <form onSubmit={handleCreatePost} className={`sm:border-x ${posts.length == 0 && "border-b"} border-white/22 relative cursor-pointer p-4 flex flex-col space-y-4`}>
                    <textarea value={content} onChange={(e) => setContent(e.target.value)} placeholder="What's happening?" className='border-none outline-none' rows={2} />
                    <label htmlFor='upload-image'>
                        <div className='p-2 w-8 -mb-6 cursor-pointer hover:bg-blue-400/20 rounded-full text-blue-400 duration-100'>
                            <Image className='w-4 h-4' />
                        </div>
                    </label>
                    <input id='upload-image' type='file' multiple accept='image/*' onChange={(e) => {
                        if (!e.target.files) return
                        setFiles(Array.from(e.target.files))
                    }} className='hidden' />
                    <button type='submit' className='bg-white cursor-pointer text-black hover:bg-white/80 duration-100 font-bold w-20 rounded-full absolute right-4 bottom-0 p-1'>
                        Post
                    </button>
                    <div className='text-sm text-red-500'>
                        {errorCreate}
                    </div>
                </form>
                <div className=''>
                    {
                        posts.map((post, index) => {
                            const idxCondition = index == posts.length - 1
                            return (
                                <Posts post={post} setPosts={setPosts} idxCondition={idxCondition} postCommentsPage={() => navigate(`/${post.id}/comments`)} onPostComment={false} setImg={setImg} />
                            )
                        })
                    }
                </div>
            </div>
        </>
    )

}

export default Feed