var gulp = require('gulp');
var sass = require('gulp-sass');
var browserify = require("browserify");
var babelify = require("babelify");
var uglify = require('gulp-uglify');
var sourcemaps = require('gulp-sourcemaps');
var browserify = require('browserify');
var source = require('vinyl-source-stream');
var buffer = require('vinyl-buffer');
var gutil = require('gulp-util');

gulp.task('default', ['sass', 'js', 'watch']);
 
gulp.task('sass', function () {
  gulp.src('./scss/**/*.scss')
    .pipe(sourcemaps.init())
    .pipe(sass().on('error', sass.logError))
    .pipe(sourcemaps.write())
    .pipe(gulp.dest('./dist/css'));
});

gulp.task('js', function() {
  var b = browserify({
      entries: './js/sg.js',
      insertGlobals: true,
      debug: true,
      transform: [babelify]
    });

  return b.bundle()
    .pipe(source('app.js'))
    .pipe(buffer())
    .pipe(sourcemaps.init({loadMaps: true}))
    // .pipe(uglify())
    .on('error', gutil.log)
    .pipe(sourcemaps.write('./'))
    .pipe(gulp.dest('./dist/js/'));

});
 
gulp.task('watch', function () {
  gulp.watch('./scss/**/*.scss', ['sass']);
  gulp.watch('./js/**/*.js', ['js']);
});