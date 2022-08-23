export enum TaskPriority{
    CRITICAL = 0,
    URGENT = 1,
    IMPORTANT = 2,
    NORMAL = 3,
    EVENTUALLY = 4,
}

export interface TodoItem {
    todoID : number
    name : string
    description : string
    priority : TaskPriority
    done : boolean
    deadline : Date
    createdAt : Date
}

export interface TodoEditor {
    name : string
    description : string
    priority : TaskPriority
    done : boolean
    deadline : Date
}

export interface TasksResponse {
    todoItemList : TodoItem[]
}

export interface LoginInfo {
    username : string
    password : string
}

export interface RegistrationInfo {
    username : string
    email : string
    password : string
}

export interface SessionIdResponse{
    sessionID : string
}