import { Skeleton } from '@mui/material'

const LoadingSkeleton = () => {
    return (
        <>
            <div className='flex flex-col gap-2'>
                <div className='flex items-center gap-2'>
                    <Skeleton sx={{ bgcolor: 'rgba(255,255,255,0.12)' }} variant="circular" width={40} height={40} />
                    <div className='flex flex-col gap-1'>
                        <Skeleton sx={{ bgcolor: 'rgba(255,255,255,0.12)' }} variant="rounded" width={120} height={10} />
                        <Skeleton sx={{ bgcolor: 'rgba(255,255,255,0.12)' }} variant="rounded" width={120} height={10} />
                    </div>
                </div>
                <Skeleton sx={{ bgcolor: 'rgba(255,255,255,0.12)' }} variant="rectangular" className='w-full' height={118} />
                <Skeleton sx={{ bgcolor: 'rgba(255,255,255,0.12)' }} variant="rectangular" className='w-full' />
                <Skeleton sx={{ bgcolor: 'rgba(255,255,255,0.12)' }} variant="rectangular" className='w-full' />
            </div>
        </>
    )
}

export default LoadingSkeleton