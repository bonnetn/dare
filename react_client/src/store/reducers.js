const EMPTY = {
    loading: false,
    tasks: [],
};

export const rootReducer = (state = EMPTY, action) => {
    switch (action.type) {
        case 'SET_LOADING':
            return {
                ...state,
                loading: action.state,
            };

        case 'SET_TASKS':
            const tasks = action.tasks.map(
                ({uuid, name, content, version}) => ({
                    uuid: uuid,
                    name: name,
                    content: content,
                    version: version,
                })
            );
            return {
                ...state,
                tasks: tasks,
            };

        default:
            return state
    }
};

