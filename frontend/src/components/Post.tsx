import { Ellipsis, Trash } from 'lucide-react'
import { useEffect, useRef, useState } from 'react'
import useDeletePosts from '../hooks/useDeletePosts'
import type { PostsWithUsername } from '../hooks/useGetAllPosts'
import defaultProfile from '../assets/defaultProfile.jpg'
import { IoChatbubbleOutline, IoHeartOutline, IoHeartSharp } from 'react-icons/io5'
import { useNavigate } from 'react-router-dom'
import useLikes from '../hooks/useLikes'
import useGetMe from '../hooks/useGetMe'


const Post = ({ post, setPosts, setPost, idxCondition, postCommentsPage, onPostComment, setImg }: { post: PostsWithUsername, setPosts?: React.Dispatch<React.SetStateAction<PostsWithUsername[]>>, setPost?: React.Dispatch<React.SetStateAction<PostsWithUsername>>, idxCondition: boolean, postCommentsPage: () => void, onPostComment: boolean, setImg: React.Dispatch<React.SetStateAction<string>> }) => {
    const [isLike, setIsLike] = useState(post.is_liked)
    const [likesCount, setLikesCount] = useState(post.likes_count)
    const [isComment, setIsComment] = useState(false)
    const [open, setOpen] = useState<number | null>(null)
    const HOST = import.meta.env.VITE_API_HOST;
    const { deletePost } = useDeletePosts(setPosts)
    const { likes } = useLikes({ setPosts, setPost })
    const { Me } = useGetMe()

    const navigate = useNavigate()

    const menuRef = useRef<HTMLDivElement | null>(null)
    useEffect(() => {
        const handleClickOutside = (e: MouseEvent) => {
            if (menuRef.current && !menuRef.current.contains(e.target as Node)) {
                setOpen(null)
            }
        }

        document.addEventListener("mousedown", handleClickOutside)
        return () => {
            document.removeEventListener("mousedown", handleClickOutside)
        }
    }, [])

    return (
        <>
            <div onClick={postCommentsPage} className={`border-t p-4 sm:border-l sm:border-r ${idxCondition && "border-b"} border-white/22 cursor-pointer relative`} key={post.id}>
                <div className='flex justify-between'>
                    <div className='flex gap-2'>
                        <img src={defaultProfile} alt="default" className='w-8 rounded-full' />
                        <div className='flex gap-2 items-center'>
                            <div className='font-bold'>
                                {post?.first_name + " " + post?.last_name}
                            </div>
                            <div className='text-gray-400'>
                                @{post?.username}
                            </div>
                            <div className='text-gray-400 text-xs'>
                                {post.updated_at.split("T")[0]}
                            </div>

                        </div>
                    </div>
                    <button onClick={(e) => {
                        e.stopPropagation()
                        setOpen(prev => prev == post.id ? null : post.id)
                    }} className='cursor-pointer hover:bg-gray-900 px-1 rounded-lg'>
                        <Ellipsis />
                    </button>
                </div>

                <div className='my-4'>
                    {post.content}
                </div>
                <div className={`grid ${post.image_url && post.image_url.length > 1 ? "grid-cols-2" : "grid-cols-1"} gap-1`}>

                    {
                        post.image_url && post.image_url.map((image, index) => (
                            <img onClick={(e) => {
                                e.stopPropagation()
                                setImg(image)
                            }} key={index} className={`
                            w-full 
                             ${post.image_url.length == 4 && index == 0 ? "rounded-tl-lg" : post.image_url.length == 4 && index == 1 ? "rounded-tr-lg" : post.image_url.length == 4 && index == 2 ? "rounded-bl-lg" : post.image_url.length == 4 && index == 3 && "rounded-br-lg"} 
                             ${post.image_url.length == 3 && index == 0 ? "h-full sm:h-full row-span-2 rounded-l-lg" : post.image_url.length == 3 && index == 1 ? "rounded-tr-lg" : post.image_url.length == 3 && index == 2 && "rounded-br-lg"}
                             ${post.image_url.length == 2 && index == 0 ? "rounded-l-lg" : post.image_url.length == 2 && index == 1 && "rounded-r-lg"}
                             ${post.image_url.length == 1 ? "h-full rounded-lg" : "h-40 sm:h-60"} 
                                object-cover
                            `}
                                src={`${HOST}/${image}`}
                            />
                        ))
                    }
                </div>

                <div className='flex justify-around space-x-4 mt-2'>
                    <button onClick={() => setIsComment(!isComment)} className='cursor-pointer gap-1'>
                        <div className='flex gap-1 hover:text-blue-400 text-xs duration-200 items-center'>
                            <div className='hover:bg-blue-300/20 rounded-full p-2 duration-200'>
                                <IoChatbubbleOutline className='w-4 h-4' />
                            </div>
                            <span className='-ml-1'>
                                {post.comments_count}
                            </span>
                        </div>
                    </button>
                    <button onClick={(e) => {
                        e.stopPropagation()
                        setIsLike(!isLike)
                        setLikesCount((prev) => {
                            if (isLike) {
                                return prev - 1
                            } else {
                                return prev + 1
                            }
                        })
                    }} className='cursor-pointer'>
                        <div className='flex gap-1 hover:text-pink-400 text-xs duration-200 items-center'>
                            <button onClick={() => likes(post.id, "POST")} className='hover:bg-pink-300/20 rounded-full p-2 duration-200 cursor-pointer'>
                                {isLike ? <IoHeartSharp className='w-4 h-4 text-pink-600' /> : <IoHeartOutline className='w-4 h-4' />}
                            </button>
                            <span className='-ml-1'>
                                {likesCount}
                            </span>
                        </div>
                    </button>

                </div>
                {
                    open == post.id && (
                        <div ref={menuRef} className='bg-black border border-white/30 p-1 rounded-lg absolute right-0 top-10'>
                            {
                                Me.id == post.user_id ? (
                                    <>
                                        <button onClick={async (e) => {
                                            e.stopPropagation()
                                            deletePost(post.id)

                                            if (onPostComment) {
                                                navigate(-1)
                                            }
                                        }} className='flex gap-2 hover:bg-white/10 px-2 rounded-sm w-full cursor-pointer text-red-500'>
                                            <Trash className='w-4' />
                                            <div>
                                                delete
                                            </div>
                                        </button>

                                    </>
                                ) : (
                                    <>
                                        <button className='flex gap-2 hover:bg-white/10 px-2 rounded-sm w-full cursor-pointer text-blue-400'>
                                            <div>
                                                (Work in Progress)
                                            </div>
                                        </button>
                                    </>
                                )
                            }
                        </div>
                    )
                }
            </div >
        </>
    )

}

export default Post