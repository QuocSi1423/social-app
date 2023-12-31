import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import {BiSolidHeart, BiHeart, BiCommentDetail, BiSolidSend} from "react-icons/bi"
import BriefUserInformation from "./BriefUserInformation";
import { getBriefUserInformation } from "../service/service-user";
import { checkInteractOfPost, getAPost, likeHandler } from "../service/service-post";
import { getCookie } from "../common";
const Post = (props) =>
{
    function formatRelativeTime(datetimeString) {
        let apiDate = new Date(datetimeString);
        let currentDate = new Date();

        let timeDifferenceInSeconds = Math.floor((currentDate - apiDate) / 1000);

        if (timeDifferenceInSeconds < 60) {
            return "Gần đây";
        } else if (timeDifferenceInSeconds < 120) {
            return "1 phút trước";
        } else if (timeDifferenceInSeconds < 3600) {
            let minutesAgo = Math.floor(timeDifferenceInSeconds / 60);
            return minutesAgo + " phút trước";
        } else if (timeDifferenceInSeconds < 7200) {
            return "1 giờ trước";
        } else if (timeDifferenceInSeconds < 86400) {
            let hoursAgo = Math.floor(timeDifferenceInSeconds / 3600);
            return hoursAgo + " giờ trước";
        } else if (timeDifferenceInSeconds < 172800) {
            return "Hôm qua";
        } else if (timeDifferenceInSeconds < 2592000) {
            let daysAgo = Math.floor(timeDifferenceInSeconds / 86400);
            return daysAgo + " ngày trước";
        } else if (timeDifferenceInSeconds < 5184000) {
            return "1 tháng trước";
        } else if (timeDifferenceInSeconds < 31536000) {
            let monthsAgo = Math.floor(timeDifferenceInSeconds / 2592000);
            return monthsAgo + " tháng trước";
        } else {
            let yearsAgo = Math.floor(timeDifferenceInSeconds / 31536000);
            return yearsAgo + " năm trước";
        }
    }


    const { postId } = props;
    const [ post, setPost ] = useState( {} );
    const [ interacted, setInteracted ] = useState( false );
    const [ user, setUser ] = useState( {} );
    const [ commentDisplay, setCommentDisplay ] = useState( false );
    const [ comment, setComment ] = useState( "" );

    useEffect( () =>
    {
        getAPost( postId ).then( result =>
        {
            setPost( result.data.data );
            getBriefUserInformation( result.data.data.user_id ).then( ( result ) =>
            {
                setUser( result.data.data );
            })
        } )
        checkInteractOfPost( postId, getCookie("accessToken") ).then( result =>{
            
                setInteracted(true)
            } )
            .catch( (err) =>
            {
                
            })
            
        }, [] )
    const interactHandle = () =>
    {
        let interactChange = !interacted ? 1 : -1;
        if ( !interacted )
        {
            likeHandler( postId ).then( r =>
            {
                console.log(r)
            } ).catch( err =>
            {
                console.log(err)
            })
        }
        setPost( { ...post, count_interact: post.count_interact + interactChange } );
        setInteracted( !interacted );
    }
    return (
        <div className="post tag">
            <BriefUserInformation avatarUrl={ user.avatar_image_url } username={ user.user_name } />
            <div className="post-content">
                <span className="post-content-time">{formatRelativeTime(post.create_at)}</span>
                <h3 className="post-content-title">{ post.title }</h3>
                <p className="post-content-des">{ post.description }</p>
                <a className="post-content-img" href={post.image_url} target="_blank">
                    <img src={ post.image_url } alt="" />
                </a>
            </div>
            <div className="post-interact">
                <button onClick={ () => { interactHandle() } }>
                    {
                        interacted?<BiSolidHeart size="30" />:<BiHeart size="30" />
                    }
                    { post.count_interact }
                </button>
                <button><BiCommentDetail size="30" onClick={()=>setCommentDisplay(true)} />{post.count_comment}</button>
            </div>
            {
                commentDisplay ?
                <div className="post-comment">
                    <input type="text" value={comment} placeholder="Comment..." onChange={(e)=>setComment(e.target.value)} />
                    <button className="post-comment-send" >
                        <BiSolidSend size={30}/>
                    </button>
                </div>
                :
                <></>
            }
            
            
        </div>
    )
}

export default Post