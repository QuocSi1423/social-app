import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import BriefUserInformation from "../components/BriefUserInformation";
import { useSelector } from "react-redux/es/hooks/useSelector";
import "../styles/views/Home.scss";
import Post from "../components/Post";
import { getBriefUserInformation } from "../service/service-user";
import { getAPost, getUserSPosts } from "../service/service-post";
import { getCookie } from "../common";
const Home = () =>
{
    const [ userInformation, setUserInformation ] = useState( {} );
    const [ posts, setPosts ] = useState( [ "conan-1", "conan-2", "85bcca41-c5aa-43db-9e44-564f3b2f4bb5" ] );
    
    

    useEffect( () =>
    {
        let id = getCookie("id")
        getBriefUserInformation(id).then( (result) =>
        {
            setUserInformation( result.data.data );
        } ).catch( err =>
        {
            console.log(err)
        })
    },[])
    return (
        <div className="home">
            <div className="home-user tag">
                <Link to="/me">
                    <img className="briefUserInformation-image" src={ userInformation.avatar_image_url } alt="" />
                </Link>
                <button className="home-user-posting">{userInformation.user_name}, write something</button>
            </div>
            <div className="home-switch">
                <button className="home-switch-follow">Theo dõi</button>
                <button className="home-switch-explore">Khám phá</button>
            </div>
            <div className="home-posts">
                {
                    posts.length > 0 ?
                        posts.map( post =>
                        {
                            return (
                                <Post key={post} postId={post} />
                            )
                            
                        } )
                        :
                        <></>
                }
                
            </div>
        </div>
    )
}

export default Home