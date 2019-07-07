import {v4} from 'uuid';


export const setLoading = (state) => ({
    type: 'SET_LOADING',
    state: state,
});

export const setTasks = (tasks) => ({
    type: 'SET_TASKS',
    tasks: tasks,
});

const proto = require('../gen/dare_grpc_web_pb.js');
const client = new proto.TaskServicePromiseClient('http://127.0.0.1:8080');

function mapProtobufTaskToJS(task) {
    return {
        uuid: task.getUuid(),
        version: task.getVersion(),
        name: task.getName(),
        content: task.getContent(),
    }
}

export function deleteTask(uuid) {
    return async function (dispatch) {
        const request = new proto.DeleteRequest();
        request.setUuid(uuid);

        await client.delete(request, {});
        dispatch(fetchTasks());
    }
}

export function editTask(uuid, version, name, content) {
    return async function (dispatch) {
        const task = new proto.Task();
        task.setUuid(uuid);
        task.setVersion(version + 1);
        task.setName(name);
        task.setContent(content);

        const request = new proto.UpsertRequest();
        request.setTask(task);

        await client.upsert(request, {
            'request-uuid': v4(),
        });
        dispatch(fetchTasks());
    }
}

export function addTask(name, content) {
    return async function (dispatch) {
        const task = new proto.Task();
        task.setName(name);
        task.setContent(content);

        const request = new proto.UpsertRequest();
        request.setTask(task);

        await client.upsert(request, {
            'request-uuid': v4(),
        });
        dispatch(fetchTasks());
    }
}

export function fetchTasks() {
    return async function (dispatch) {
        dispatch(setLoading(true));

        const request = new proto.GetAllRequest();
        const response = await client.getAll(request, {});
        const protobufTaskList = response.getTasksList();
        const taskList = protobufTaskList.map(mapProtobufTaskToJS);

        dispatch(setTasks(taskList));
        dispatch(setLoading(false));
    }
}

