import React from 'react';
import './App.css';
import {fetchTasks} from "./store/actions";
import {connect} from "react-redux";
import TaskList from "./components/task-list/task-list";

class App extends React.Component {
    render() {
        return (
            <div>
                Hello
                <button onClick={this.props.onClick}>GO</button>
                <TaskList/>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return {
        tasks: state.tasks,
        loading: state.loading,
    };
};
const mapDispatchToProps = (dispatch) => {
    return {
        onClick: () => {
            dispatch(fetchTasks())
        }
    }
};

export default connect(
    mapStateToProps,
    mapDispatchToProps,
)(App);

//export default App;
