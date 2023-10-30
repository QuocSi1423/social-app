import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import {BiSolidHeart, BiHeart, BiCommentDetail} from "react-icons/bi"
import BriefUserInformation from "./BriefUserInformation";
import { getBriefUserInformation } from "../service/service-user";
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


    const { post } = props;
    const [ interacted, setInteracted ] = useState( false );
    const [ user, setUser ] = useState( {} );

    useEffect( () =>
    {
        getBriefUserInformation( post.user_id ).then( ( result ) =>
        {
            setUser( result.data.data );
        })
    }, [] )
    const interactHandle = () =>
    {
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
                <button><BiCommentDetail size="30" />{post.count_comment}</button>
            </div>
            
        </div>
    )
}

export default Post