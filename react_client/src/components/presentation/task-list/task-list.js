import React from 'react';
import TaskDisplay from "../task-display/task-display";
import PropTypes from 'prop-types';


class TaskList extends React.Component {
    render() {
        const taskList = this.props.tasks.map((task) =>
            <li key={task.uuid}>
                <TaskDisplay task={task}/>
            </li>
        );
        return (
            <ul>
                {taskList}
            </ul>
        );
    }
}

TaskList.propTypes = {
    tasks: PropTypes.arrayOf(PropTypes.shape({
        uuid: PropTypes.string.isRequired,
    })).isRequired,
};

export default TaskList;

