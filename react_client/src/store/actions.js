export const setLoading = (state) => ({
    type: 'SET_LOADING',
    state: state,
});

export const setTasks = (tasks) => ({
    type: 'SET_TASKS',
    tasks: tasks.map(({uuid, name, content}) => ({
        uuid: uuid,
        name: name,
        content: content,
    })),
});

const proto = require('../gen/dare_grpc_web_pb.js');
const client = new proto.TaskServicePromiseClient('http://127.0.0.1:8080');

function mapProtobufTaskToJS(task) {
    return {
        uuid: task.getUuid(),
        name: task.getName(),
        content: task.getContent(),
    }
}

export function fetchTasks() {
    // Thunk middleware knows how to handle functions.
    // It passes the dispatch method as an argument to the function,
    // thus making it able to dispatch actions itself.
    return async function (dispatch) {
        // First dispatch: the app state is updated to inform
        // that the API call is starting.

        dispatch(setLoading(true));

        // The function called by the thunk middleware can return a value,
        // that is passed on as the return value of the dispatch method.

        // In this case, we return a promise to wait for.
        // This is not required by thunk middleware, but it is convenient for us.

        const request = new proto.GetAllRequest();
        const response = await client.getAll(request, {});
        const protobufTaskList = response.getTasksList();
        const taskList = protobufTaskList.map(mapProtobufTaskToJS);

        dispatch(setTasks(taskList));
        dispatch(setLoading(false));

    }
}

