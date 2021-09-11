<template>
  <div id="app" class="container-fluid text-center">
    <h1 class="text-info">{{ title }}</h1>
    <br />

    <div class="col">
      <div v-if="error" class="alert alert-danger error" @click="error = !error">
        <strong>Error:</strong> Lütfen Bir Task Giriniz.
      </div>
      <form @submit.prevent="addTask">
        <div class="input-group mb-3">
          <input
            type="text"
            class="form-control"
            placeholder="Task Name"
            v-model="taskName"
          />
          <div class="input-group-append">
            <button class="btn btn-success" type="submit">
              <i class="fa fa-plus" aria-hidden="true"></i>
            </button>
          </div>
        </div>
      </form>
      <ul class="list-group">
        <li
          v-for="(task_name, index) in tasks"
          :key="index"
          class="list-group-item list-group-item-info"
        >
          <div class="row">
            <div class="col">
              {{ task_name["task"] }}
            </div>
            <div class="col-1">
              <a>
                <button
                  type="button"
                  class="btn btn-danger"
                  @click="deleteToDoItem(task_name['id'])"
                >
                  Sil
                </button>
              </a>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "App",
  data() {
    return {
      title: "ToDo List",
      taskName: "",
      tasks: [],
      error: false,
    };
  },
  created() {
    axios.get("http://localhost:8081/getItems").then((resp) => {
      let resTasks = Object.keys(resp.data).map((key) => {
        return {
          id: key,
          task: resp.data[key].item,
        };
      });
      this.tasks = resTasks;
    });
  },
  methods: {
    deleteToDoItem: function (itemId) {
      axios
        .post("http://localhost:8081/deleteItem?removeID=" + itemId)
        .then(() => {
          this.tasks.splice(
            this.tasks.indexOf(this.tasks.find((item) => item.id === itemId)),
            1
          );
        });
    },
    addTask: function () {
      if (this.taskName === '') {
        this.error = true;
        return;
      } else {
        this.error = false;
      }

      const postTask = { item: this.taskName };
      axios.post("http://localhost:8081/postItem", postTask).then((resp) => {
        this.tasks.push({ id: resp.data.name, task: postTask.item });
      });
      this.taskName = "";
    },
  },
};
</script>
