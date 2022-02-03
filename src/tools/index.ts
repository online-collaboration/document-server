import { nanoid } from 'nanoid'

/**
 * 获取唯一ID
 * @returns 
 */
export const getUniqueID = ()=> {
    return nanoid()
}