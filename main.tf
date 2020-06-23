provider "dcos" { 
  dcos_url  = "http://gdmst001v.gsil.rri-usa.org"
  user = "gsil-readonly"
  password = "Gmf5ZF7k9LYJTmpwKMeG"
}


resource "dcos_ham_app" "my-server" {
  app_id = "/gsiltest/terraform-nginx-1"
  cpus      = 0.1
  mem       = 32
  instances = 1

  labels  = {
    version = "1.0.0"
  }
  
  container {
    type    = "DOCKER"

    docker {
      image = "nginx:stable-alpine" 
    }
    
    volumes {
        host_path       = "/nfs/app/common/test_html/index_1.1.html" 
        container_path  = "/usr/share/nginx/html/index.html"
        mode            = "RW"
    }
  }  

  gsil_filename = "done.zip"
  gsil_directory = "foo"
  gsil_which_poke = 100

  gsil_injector {
    username = "ggaspar"
    password = var.injector_password
    location = "http://gdmst001v.gsil.rri-usa.org/"

    entries {
      entry {
        key_name = "A_KEYFILE_NAME"
        value = "bar"
        path = "/gsiltest/autofix"
        secret = false
        function = "PUT"
      }
      entry {
        key_name = "B_KEYFILE_NAME_TWO"
        value = "bazz"
        path = "/gsiltest/autofix"
        secret = false
        function = "PUT"
      }
    }
  }

}  

