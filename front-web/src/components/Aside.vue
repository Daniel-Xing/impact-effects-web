<template>
  <div v-loading="loading" style="height: 100%;">
    <el-menu ref="menu" :router="true" :default-active="activedIndex" style="height: 100%;">
      <el-menu-item :route="{ path: '/dashboard' }" index="/dashboard">
        <template slot="title">
          <i class="el-icon-pie-chart"></i>
          <span>Discription</span>
        </template>
      </el-menu-item>
      <!-- <el-menu-item v-if="account.group == 3" :route="{ path: '/user' }" index="/user">
        <template slot="title">
          <i class="el-icon-user"></i>
          <span>用户管理</span>
        </template>
      </el-menu-item>
      <el-menu-item v-if="account.group == 3" :route="{ path: '/label' }" index="/label">
        <template slot="title">
          <i class="el-icon-camera"></i>
          <span>图片标签</span>
        </template>
      </el-menu-item>
      <el-menu-item v-if="account.group == 3" :route="{ path: '/videolabel' }" index="/videolabel">
        <template slot="title">
          <i class="el-icon-video-camera"></i>
          <span>视频标签</span>
        </template>
      </el-menu-item>
      <el-menu-item v-if="account.group == 3" :route="{ path: '/partition' }" index="/partition">
        <template slot="title">
          <i class="el-icon-finished"></i>
          <span>任务设置</span>
        </template>
      </el-menu-item> -->
      <el-menu-item :route="{ path: '/list' }" index="/list">
        <template slot="title">
          <i class="el-icon-document-checked"></i>
          <span>Single Impactor</span>
        </template>
      </el-menu-item>
    </el-menu>
  </div>
</template>
<script>
import { getTaskList as getList } from "@/models/Task.js";
export default {
  name: "Aside",
  data() {
    return {
      isCollapse: false,
      taskList: this.global.task,
      activedIndex: this.$route.path,
      account: this.global.account,
      loading: false,
    };
  },
  methods: {
    // getTaskList: function() {
    //   if (this.account.group == 0) {
    //     return false;
    //   }
    //   this.loading = true;
    //   getList(this.account)
    //     .then(data => {
    //       this.global.task.list = data.data.taskList ? data.data.taskList : data.data;
    //     })
    //     .catch(data => {
    //       this.$message({
    //         type: "error",
    //         message: "获取任务列表失败"
    //       });
    //     })
    //     .then(() => {
    //       this.loading = false;
    //     });
    // }
  },
  watch: {
    "$route.path": function () {
      this.activedIndex = this.$route.path;
    }
  },
  mounted: function () {
    // this.getTaskList();
  }
};
</script>
<style scoped>
.task-tag {
  margin: -3px 0 0 5px;
}
</style>
