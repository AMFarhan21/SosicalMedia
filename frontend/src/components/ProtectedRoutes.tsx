import type { JSX } from "react"
import { Navigate } from "react-router-dom"
import useGetMe from "../hooks/useGetMe"

const ProtectedRoutes = ({ children }: { children: JSX.Element }) => {

    const { errorGetMe, loadingGetMe } = useGetMe()
    const token = localStorage.getItem("Token")

    if (loadingGetMe) return null

    if (!token) {
        return <Navigate to="/auth" replace />
    }

    if (errorGetMe === "UNAUTHORIZED") {
        localStorage.removeItem("Token")
        return <Navigate to="/auth" replace />
    }

    return children
}

export default ProtectedRoutes