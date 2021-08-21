import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'
import Constants from '../constants'

import currentUserReducer from './currentUserSlice'

export function makeStore() {
  return configureStore({
    reducer: {
      currentUser: currentUserReducer,
    },
    devTools: !Constants.IS_LIVE
  })
}

const store = makeStore()

export type AppState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  AppState,
  unknown,
  Action<string>
>

export default store