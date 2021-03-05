
#### 运行
```shell
docker run -it -d -p 4869:4869 -v /data/zimg/:/zimg/bin/img --name zimg 52334755/zimg
```


#### Javascript

```vue
<template>
  <div>
    <el-upload
      action="http://localhost:4869/upload"
      :show-file-list="false"
      :http-request="uploadImg"
      :on-success="onSuccess"
      :on-error="onError"
      :on-remove="onRemove"
    >
      <el-button size="small" type="primary">选择图片</el-button>
    </el-upload>
  </div>
</template>

<script>
export default {
  components: {},
  data() {
    return {};
  },
  created() {},
  methods: {
    onRemove(file, fileList) {
      console.log("onRemove", file);
    },
    onSuccess(res, file) {
      console.log("onSuccess", res, file);
    },
    onError() {
      console.log("onError", arguments);
    },
    async uploadImg(f) {
      const { data: resp } = await this.$axios.post("http://localhost:4869/upload", f.file, {
        headers: { "content-type": "jpeg" },
      });
      if (resp.ret) {
        console.log("Upload success, md5:", resp.info.md5);
        console.log("http://localhost:4869/"+ resp.info.md5);
      } else {
        console.log("Upload error:", resp.error.message);
      }
    },
  },
};
</script>
```

```golang


```
