<template>
  <div class="cmdb-search-container">
    <!-- 搜索框 -->
    <div class="box-shadow search-box">
      <div class="search-input">
        <el-input
          v-model="query.keywords"
          placeholder="请输入 资源名称|ID|IP 敲回车进行搜索"
          @keyup.enter.native="search"
        ></el-input>
      </div>
    </div>

    <!-- 搜索结果 -->
    <div class="content-box">
      <el-table
        :data="resources"
        v-loading="fetchResourceLoading"
        style="width: 100%"
      >
        <el-table-column prop="name" label="名称">
          <template slot-scope="{ row }">
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="资产来源">
          <template slot-scope="{ row }">
            {{ row.vendor }} <br />
            {{ row.region }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="内网IP/外网IP">
          <template slot-scope="{ row }">
            {{ row.private_ip }} <br />
            {{ row.public_ip }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="同步时间">
          <template slot-scope="{ row }">
            {{ row.sync_at | parseTime }}
          </template>
        </el-table-column>
        <el-table-column prop="sync_at" label="创建时间">
          <template slot-scope="scope">
            {{ scope.row.create_at | parseTime }}
          </template>
        </el-table-column>
        <el-table-column prop="expire_at" label="过期时间">
          <template slot-scope="scope">
            {{ scope.row.expire_at | parseTime }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="状态">
          <template slot-scope="{ row }">
            {{ row.status }}
          </template>
        </el-table-column>
      </el-table>

      <pagination
        v-show="total > 0"
        :total="total"
        :page.sync="query.page_number"
        :limit.sync="query.page_size"
        @pagination="search"
      />
    </div>
  </div>
</template>

<script>
import { SEARCH } from "@/api/cmdb/resource";
import Pagination from "@/components/Pagination";

export default {
  name: "Search",
  components: { Pagination },
  data() {
    return {
      fetchResourceLoading: false,
      resources: [],
      total: 0,
      query: {
        page_size: 20,
        page_number: 1,
        keywords: "",
      },
    };
  },
  methods: {
    async search() {
      this.fetchResourceLoading = true;
      try {
        const resp = await SEARCH(this.query);
        this.resources = resp.data.items;
        this.total = resp.data.total;
      } catch (error) {
        this.$notify.error({
          title: "搜索资源异常",
          message: error,
        });
      } finally {
        this.fetchResourceLoading = false;
      }
    },
  },
};
</script>

<style scoped>
.search-box {
  height: 40px;
  background-color: white;
  display: flex;
  align-items: center;
  border-radius: 4px;
}

.search-input {
  width: 100%;
  margin-left: 12px;
  margin-right: 12px;
}

.content-box {
  padding-top: 12px;
}
</style>
