def require_dir(dirname)
  Dir[File.join(File.dirname(__FILE__), dirname, '**', '*.rb')].sort.each {|file| require file }
end