<template>
  <div>
    <tips :tips="tips" />

    <!-- 表格功能区 -->
    <div class="table-op">
        <div class="search">
        <el-input
            v-model="query.keywords"
            placeholder="请输入凭证描述|API KEY|用户名 敲回车进行搜索"
            @keyup.enter.native="get_secrets"
        ></el-input>
        </div>
        <div class="op">
        <el-button icon="el-icon-plus" type="primary" @click="handleAddSecret">添加凭证</el-button>
        </div>
    </div>

    <!-- 凭证表格 -->
    <div class="box-shadow secret-box">
      <el-table
        :data="secrets"
        v-loading="fetchSecretLoading"
        style="width: 100%"
      >
        <el-table-column prop="name" label="描述">
          <template slot-scope="{ row }">
            {{ row.description }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="厂商">
          <template slot-scope="{ row }">
            {{ row.vendor }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="类型">
          <template slot-scope="{ row }">
            {{ row.crendential_type }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="凭证">
          <template slot-scope="{ row }">
            {{ row.api_key }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="同步地域">
          <template slot-scope="{ row }">
            {{ row.allow_regions.join(",") }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="创建时间">
          <template slot-scope="{ row }">
            {{ row.create_at | parseTime }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="速率限制">
          <template slot-scope="{ row }">
            {{ row.request_rate }}
          </template>
        </el-table-column>
        <el-table-column prop="操作" align="center" label="状态">
          <template>
            <el-button type="text" disabled>删除</el-button>
            <el-button type="text" disabled>禁用</el-button>
            <el-button type="text" disabled>测试</el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination
        v-show="total > 0"
        :total="total"
        :page.sync="query.page_number"
        :limit.sync="query.page_size"
        @pagination="get_secrets"
      />
    </div>

    <!-- 添加secret -->
    <add-secret :visible.sync="showAddSecretDrawer" />
  </div>
</template>

<script>
import Tips from "@/components/Tips";
import Pagination from "@/components/Pagination";
import { LIST_SECRET } from "@/api/cmdb/secret";
import AddSecret from "./components/AddSecret";

export default {
  name: "ResourceSync",
  components: { Tips, Pagination, AddSecret },
  data() {
    return {
      secrets: [],
      fetchSecretLoading: false,
      total: 0,
      query: {
        page_size: 20,
        page_number: 1,
      },
      showAddSecretDrawer: false,
    };
  },
  mounted() {
    this.get_secrets();
  },
  methods: {
    async get_secrets() {
      this.fetchSecretLoading = true;
      try {
        const resp = await LIST_SECRET(this.query);
        this.secrets = resp.data.items;
        this.total = resp.data.total;
      } catch (error) {
        this.$notify.error({
          title: "获取凭证异常",
          message: error,
        });
      } finally {
        this.fetchSecretLoading = false;
      }
    },
    handleAddSecret() {
      this.showAddSecretDrawer = true;
    },
  },
};
</script>

<style scoped>
.secret-box {
  margin-top: 8px;
}

.table-op {
  margin-top: 12px;
  display: flex;
  align-items: center;
}

.op {
    margin-left: auto;
}
</style>