task :default do
  sh 'rake -T'
end

namespace :prepare do
  desc 'go'
  task :go do
    sh 'go get -u github.com/labstack/echo'
    sh 'go get -u github.com/facebookgo/grace/gracehttp'
  end

  desc 'glide install'
  task :glide do
    sh 'glide install'
  end

  desc 'gems'
  task :gems do
    sh 'gem install guard guard-shell'
  end
end

namespace :server do
  desc 'server run'
  task :run do
    sh 'go run server.go'
  end
end

namespace :info do
  desc 'grace'
  task :grace do
    puts "how to graceful shutdown:\n\t`kill -USR2 pid`"
    sh 'xdg-open https://github.com/facebookgo/grace'
  end

  desc 'echo'
  task :echo do
    sh 'xdg-open https://github.com/facebookgo/grace'
  end
end

namespace :deploy do
  desc 'hsj'
  task :hsj do
    sh 'go build'
    sh 'scp gomonkey itang@haoshuju.net:gomonkey/'
    sh 'rm gomonkey'
  end
end
