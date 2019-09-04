# frozen_string_literal: true

require 'spec_helper'
require 'job_sorter'

RSpec.describe JobSorter do
  describe 'empty input string' do
    it 'returns an empty sequence' do
      expect(JobSorter.new('').to_s).to     be_empty
      expect(JobSorter.new(' ').to_s).to    be_empty
      expect(JobSorter.new(nil).to_s).to    be_empty
    end
  end

  describe 'a job structure of a=>, b=>, c=>  with no dependencies' do
    it 'returns a sequence containing jobs abc in no significant order' do
      string = "a=>\nb=>\nc=>"
      expectation = 'a,b,c'

      sorter = JobSorter.new(string)

      expect(sorter.to_s).to eql(expectation)
    end
  end

  describe 'a job structure of a=>, b=>c, c=> where b has a dependency' do
    it 'returns a sequence that positions c before b, and contains all three jobs abc' do
      string = "a=>\nb=>c\nc=>"

      response = JobSorter.new(string).to_s.split(',')
      index_of_b = response.index { |i| i == 'b' }
      index_of_c = response.index { |i| i == 'c' }

      expect(index_of_c).to be < index_of_b
    end
  end

  describe 'a job structure of a=>, b=>c, c=>f, d=>a, e=>b, f=>' do
    it 'returns a sequence that positions f before c, c before b, b before e, and a before d containing all six jobs abcdef' do
      string = "a=>\nb=>c\nc=>f\nd=>a\ne=>b\nf=>"

      response = JobSorter.new(string).to_s.split(',')
      index_of_a = response.index { |i| i == 'a' }
      index_of_b = response.index { |i| i == 'b' }
      index_of_c = response.index { |i| i == 'c' }
      index_of_d = response.index { |i| i == 'd' }
      index_of_e = response.index { |i| i == 'e' }
      index_of_f = response.index { |i| i == 'f' }

      expect(index_of_f).to be < index_of_c
      expect(index_of_c).to be < index_of_b
      expect(index_of_b).to be < index_of_e
      expect(index_of_a).to be < index_of_d
    end
  end

  describe 'a job structure of a=>, b=>, c=>c where a self dependency occurs' do
    it 'returns an error stating that jobs can not depend on itself' do
      string = "a=>\nb=>\nc=>c"

      expect { JobSorter.new(string) }.to raise_error(JobSorter::SelfDependencyException, 'Self dependency')
    end
  end

  describe 'a job structure of a=>, b=>c, c=>f, d=>a, e=>, f=>b where a circular dependency occurs' do
    it 'returns an error stating that jobs can not have circular dependencies' do
      string = "a=>\nb=>c\nc=>f\nd=>a\ne=>\nf=>b"

      expect { JobSorter.new(string) }.to raise_error(JobSorter::CircularDependencyException, 'Circular dependency')
    end
  end
end
