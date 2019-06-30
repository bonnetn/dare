import React from 'react';
import './task-list.css';
import {connect} from "react-redux";


class TaskList extends React.Component {
    render() {
        const taskList = this.props.tasks.map((task) =>
            <li key={task.uuid}>
                {task.name}
            </li>
        );
        return (
            <ul>
                {taskList}
            </ul>
        );
    }
}

const mapStateToProps = state => {
    return {
        tasks: state.tasks,
        loading: state.loading,
    };
};

export default connect(
    mapStateToProps,
    null,
)(TaskList);

