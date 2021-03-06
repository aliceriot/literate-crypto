export const SET_NAV_SHOW_STATE = 'SET_NAV_SHOW_STATE'
export const setNavShowState = bool => (
  { type: SET_NAV_SHOW_STATE, payload: bool }
)

export const SET_MATASANO_CONTENTS = 'SET_MATASANO_CONTENTS'
export const setMatasanoContents = contents => (
  { type: SET_MATASANO_CONTENTS, payload: contents }
)

export const SET_BLOG_CONTENTS = 'SET_BLOG_CONTENTS'
export const setBlogContents = contents => (
  { type: SET_BLOG_CONTENTS, payload: contents }
)
