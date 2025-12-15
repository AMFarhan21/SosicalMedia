import { Ellipsis, MessageCircle, SquarePen, ThumbsUp, Trash } from 'lucide-react'
import { useState } from 'react'
import useDeletePosts from '../hooks/useDeletePosts'
import type { PostsWithUsername } from '../hooks/useGetAllPosts'


const Posts = ({ posts, setPosts }: { posts: PostsWithUsername[], setPosts: React.Dispatch<React.SetStateAction<PostsWithUsername[]>> }) => {
    const [isLike, setIsLike] = useState(false)
    const [isComment, setIsComment] = useState(false)
    const [open, setOpen] = useState<number | null>(null)

    const { deletePost, errorDelete } = useDeletePosts(setPosts)



    return (
        <div className=''>
            {
                posts.map((post, index) => {
                    return (
                        <div className={`border-t p-4 sm:border-l sm:border-r ${index == posts.length - 1 && 'rounded-b-xl'} ${index == 0 && 'rounded-t-xl'} border-gray-400 cursor-pointer relative`} key={post.id}>
                            <div className='flex justify-between'>
                                <div className='flex gap-2 items-center'>
                                    <div className='font-bold'>
                                        @{post.username}
                                    </div>
                                    <div className='text-gray-400 text-xs'>
                                        {post.updated_at.split("T")[0]}
                                    </div>

                                </div>
                                <button onClick={() => setOpen(prev => prev == post.id ? null : post.id)} className='cursor-pointer hover:bg-gray-200 px-1 rounded-lg'>
                                    <Ellipsis />
                                </button>
                            </div>

                            <div className='mb-4'>
                                {post.content}
                            </div>
                            {
                                post.image_url && <img className='mb-2' src={post.image_url} />
                            }

                            <div className='flex space-x-4'>
                                <button onClick={() => setIsLike(!isLike)} className='cursor-pointer'>
                                    {
                                        isLike ? (
                                            <div className='flex gap-1'>
                                                <span>
                                                    0
                                                </span>
                                                <ThumbsUp className='w-4' />
                                                <span>
                                                    like
                                                </span>
                                            </div>
                                        ) : (
                                            <div className='text-pink-800 flex gap-1'>
                                                <span>
                                                    0
                                                </span>
                                                <ThumbsUp className='w-4' />
                                                <span>
                                                    like
                                                </span>
                                            </div>
                                        )

                                    }
                                </button>
                                <button onClick={() => setIsComment(!isComment)} className='cursor-pointer gap-1'>
                                    {
                                        isComment ? (
                                            <div className='flex gap-1'>
                                                <span>
                                                    0
                                                </span>
                                                <MessageCircle className='w-4' />
                                                comment
                                            </div>
                                        ) : (
                                            <div className='text-pink-800 flex gap-1'>
                                                <span>
                                                    0
                                                </span>
                                                <MessageCircle className='w-4' />
                                                comment
                                            </div>
                                        )

                                    }
                                </button>
                            </div>
                            {
                                open == post.id && (
                                    <div className='bg-white border p-2 rounded-lg absolute right-0 top-10 -mr-10'>
                                        <button onClick={async () => {
                                            const success = await deletePost(post.id)
                                            if (!success) {
                                                alert(errorDelete)
                                            }
                                        }} className='flex gap-2 hover:bg-gray-200 px-2 rounded-sm w-full cursor-pointer text-red-500'>
                                            <Trash className='w-4' />
                                            <div>
                                                delete
                                            </div>
                                        </button>
                                        <button className='flex gap-2 hover:bg-gray-200 px-2 rounded-sm w-full cursor-pointer text-blue-400'>
                                            <SquarePen className='w-4' />
                                            <div>
                                                edit
                                            </div>
                                        </button>
                                    </div>
                                )
                            }
                        </div>
                    )
                })
            }

        </div>
    )
}

export default Posts