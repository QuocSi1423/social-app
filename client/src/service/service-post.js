import axios from "axios"

export const getAPost = (postId) =>
{
    return (
        axios.get(`http://localhost:5173/v1/posts/${postId}`)
    )
}

export const getUserSPosts = (userId) =>
{
    return (
        axios.get(`http://localhost:5173/v1/posts/?user_id=${userId}`)
    )
}

export const getUserInteractOfPost = ( postId ) =>
{
    return axios.get(`http://localhost:5173/v1/posts/${postId}/interacts/`,{
    withCredentials: true,
})
}

export const checkInteractOfPost = ( postId, token ) =>
{
    return (
        axios.get( `http://localhost:5173/v1/posts/${ postId }/likes`,
        {headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json',
  }})
    )
}

export const likeHandler = ( postId ) =>
{
    return (
        axios.post(`http://localhost:5173/v1/posts/${postId}/likes`)
    )
}

export const incInteract = () =>
{
    
}