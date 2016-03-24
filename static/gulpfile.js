/*
 * Copyright 2016, Robert Bieber
 *
 * This file is part of senatron.
 *
 * senatron is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * senatron is distributed in the hope that it will be useful,
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with senatron.  If not, see <http://www.gnu.org/licenses/>.
 */

var babelify = require('babelify');
var browserify = require('browserify');
var fs = require('fs');
var gulp = require('gulp');
var gulpif = require('gulp-if');
var source = require('vinyl-source-stream');
var streamify = require('gulp-streamify');
var uglify = require('gulp-uglify');
var util = require('gulp-util');

var NODE_ENV = require('node-env');

gulp.task('build', function() {
	function build(fileName, i, fileNames) {
		browserify({debug: NODE_ENV !== 'production'})
			.require(
				'./js/'+fileName,
				{expose: fileName.split('.')[0]}
			).transform(babelify)
			.bundle()
			.on('error', util.log)
			.pipe(source(fileName))
			.pipe(gulpif(NODE_ENV === 'production', streamify(uglify())))
			.pipe(gulp.dest('./build/js/'));
	}

	var filePattern = /.js$/;
	fs.readdirSync(__dirname+'/js')
		.filter(filePattern.test.bind(filePattern))
		.map(build);

	gulp.src('./template/**/*.got')
		.pipe(gulp.dest('./build/template/'));

	gulp.src('./css/**/*.css')
		.pipe(gulp.dest('./build/css/'));
});

gulp.task('watch', function() {
	gulp.watch('./js/**/*.js', ['build']);
	gulp.watch('./template/**/*.got', ['build']);
	gulp.watch('./css/**/*.css', ['build']);
});

gulp.task('default', ['build', 'watch']);
