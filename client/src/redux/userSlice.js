import { createSlice } from "@reduxjs/toolkit";

const userSlice = createSlice( {
    name:"user",
    initialState: {
        loginName: ""
    },
    reducers: {
        updateUser: ( state, action ) =>
        {
            state.loginName = action.payload.loginName;
        }
    },
} )

export default userSlice.reducer;
export const { updateUser } = userSlice.actions;