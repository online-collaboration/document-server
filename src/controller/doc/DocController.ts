import { Param, Body, Get, Put, Delete, JsonController, Res } from 'routing-controllers'

export class DocController{
    /**
     * Request one doc content
     * @param id doc_id
     * @param response 
     * @returns 
     */
    @Get('/doc/:id')
    getDoc(@Param('id') id: string, @Res() response: any){
        return response.send({})
    }

    /**
     * Create one doc
     * @param id doc_id
     * @param response 
     * @returns 
     */
    @Put('/doc')
    createDoc(@Param('id') id: string, @Res() response: any){
        return 
    }

    /**
     * Delete this doc
     * @param id doc_id
     * @param response 
     * @returns 
     */
    @Delete('/doc')
    deleteDoc(@Param('id') id: string, @Res() response: any){
        return 1
    }
}