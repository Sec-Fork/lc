package cmd

const defaultConfigFile = `# # lc (list cloud) 的云服务商配置文件

# # 配置文件说明

# # provider 是云服务商的名字
# - provider: provider_name
#   # id 是当前配置文件的名字
#   id: test
#   # access_key 是这个云的访问凭证 Key 部分
#   access_key: 
#   # secret_key 是这个云的访问凭证 Secret 部分
#   secret_key: 
#   # （可选）session_token 是这个云的访问凭证 session token 部分，仅在访问凭证是临时访问配置时才需要填写这部分的内容
#   session_token: 

# # 阿里云
# # 访问凭证获取地址：https://ram.console.aliyun.com
# - provider: aliyun
#   id: aliyun_default
#   access_key: 
#   secret_key: 
#   session_token: 

# # 腾讯云
# # 访问凭证获取地址：https://console.cloud.tencent.com/cam
# - provider: tencent cloud
#   id: tencent_cloud_default
#   access_key: 
#   secret_key: 
#   session_token: 

# # 华为云
# # 访问凭证获取地址：https://console.huaweicloud.com/iam
# - provider: huawei cloud
#   id: huawei_cloud_default
#   access_key: 
#   secret_key: 
#   session_token: 

# # 天翼云
# # 访问凭证获取地址：https://oos-cn.ctyun.cn/oos/ctyun/iam/dist/index.html#/certificate
# - provider: ctyun
#  id: ctyun_default
#  access_key: 
#  secret_key:

# # 百度云
# # 访问凭证获取地址：https://console.bce.baidu.com/iam/
# - provider: baidu cloud
#   id: baidu_cloud_default
#   access_key: 
#   secret_key: 
#   session_token: 

# # 联通云
# # 访问凭证获取地址：https://console.cucloud.cn/console/uiam
# - provider: cu cloud
#   id: cu_cloud_default
#   access_key: 
#   secret_key: 
#   session_token: 
`
