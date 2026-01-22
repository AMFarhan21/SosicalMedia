import useGetMe from "../hooks/useGetMe"

const Profile = () => {
    const { Me } = useGetMe()

    const profile = ["Post", "Comments"]


    return (
        <div className='w-full sm:w-[56%] mx-auto'>
            <div className="bg-gray-600 w-full h-40 relative">
                <div className="w-30 h-30 border rounded-full bg-gray-600 bottom-0 absolute -mb-15 ml-4"></div>
            </div>
            <div className="w-full p-4">
                <div className="mt-16 font-bold text-xl">{Me.first_name + " " + Me.last_name}</div>
                <div className="text-gray-600">@{Me.username}</div>
                <div className="flex justify-around">
                    {
                        profile.map((p, key) => (
                            <div key={key}>
                                {p}
                            </div>
                        ))
                    }
                </div>
            </div>
        </div>
    )
}

export default Profile